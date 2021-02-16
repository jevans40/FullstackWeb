import 'source-map-support/register';

import type { ValidatedEventAPIGatewayProxyEvent } from '@libs/apiGateway';
import { formatJSONResponse } from '@libs/apiGateway';
import { middyfy } from '@libs/lambda';
import {promisify} from 'util';
import * as Axios from 'axios';
import * as jsonwebtoken from 'jsonwebtoken';
const jwkToPem = require('jwk-to-pem');

import schema from './schema';
import { tokenToString } from 'typescript';

/*
This is directly taken from the amazon aws github as an authentication demo
https://github.com/awslabs/aws-support-tools/blob/master/Cognito/decode-verify-jwt/decode-verify-jwt.ts
*/

interface ClaimVerifyRequest {
  readonly token?: string;
}

interface ClaimVerifyResult {
  readonly userName: string;
  readonly clientId: string;
  readonly isValid: boolean;
  readonly error?: any;
}

interface TokenHeader {
  kid: string;
  alg: string;
}
interface PublicKey {
  alg: string;
  e: string;
  kid: string;
  kty: string;
  n: string;
  use: string;
}
interface PublicKeyMeta {
  instance: PublicKey;
  pem: string;
}

interface PublicKeys {
  keys: PublicKey[];
}

interface MapOfKidToPublicKey {
  [key: string]: PublicKeyMeta;
}

interface Claim {
  token_use: string;
  auth_time: number;
  iss: string;
  exp: number;
  username: string;
  client_id: string;
}


const cognitoPoolId = process.env.COGNITO_POOL_ID || 'us-east-2_E4FNAOYCN';
if (!cognitoPoolId) {
  throw new Error('env var required for cognito pool');
}
const cognitoIssuer = `https://cognito-idp.us-east-2.amazonaws.com/${cognitoPoolId}`;

let cacheKeys: MapOfKidToPublicKey | undefined;
const getPublicKeys = async (): Promise<MapOfKidToPublicKey> => {
  if (!cacheKeys) {
    const url = `${cognitoIssuer}/.well-known/jwks.json`;
    console.log(url)
    const publicKeys = await Axios.default.get<PublicKeys>(url,{headers:{"Access-Control-Allow-Origin": "*"}});
    console.log(publicKeys)
    cacheKeys = publicKeys.data.keys.reduce((agg, current) => {
      const pem = jwkToPem(current);
      agg[current.kid] = {instance: current, pem};
      return agg;
    }, {} as MapOfKidToPublicKey);
    return cacheKeys;
  } else {
    return cacheKeys;
  }
};

const verifyPromised = promisify(jsonwebtoken.verify.bind(jsonwebtoken));

const handler = async (request: ClaimVerifyRequest): Promise<ClaimVerifyResult> => {
  let result: ClaimVerifyResult;
  let failed = 0
  let log : string[]
  try {
    console.log(`user claim verify invoked for ${JSON.stringify(request)}`);
    const token = request.token;
    const tokenSections = (token || '').split('.');
    log = tokenSections
    failed = 1
    if (tokenSections.length < 2) {
      throw new Error('requested token is invalid');
    }
    const headerJSON = Buffer.from(tokenSections[0], 'base64').toString('utf8');
    const header = JSON.parse(headerJSON) as TokenHeader;
    const keys = await getPublicKeys();
    const key = keys[header.kid];
    failed = 2
    if (key === undefined) {
      throw new Error('claim made for unknown kid');
    }
    const claim = await verifyPromised(token, key.pem) as Claim;
    const currentSeconds = Math.floor( (new Date()).valueOf() / 1000);
    failed = 3
    if (currentSeconds > claim.exp || currentSeconds < claim.auth_time) {
      throw new Error('claim is expired or invalid');
    }
    failed = 4
    if (claim.iss !== cognitoIssuer) {
      throw new Error('claim issuer is invalid');
    }
    failed = 5
    if (claim.token_use !== 'access') {
      throw new Error('claim use is not access');
    }
    failed = 6
    console.log(`claim confirmed for ${claim.username}`);
    result = {userName: claim.username, clientId: claim.client_id, isValid: true};
  } catch (error) {
    result = {userName: `${JSON.stringify(request.token)}`, clientId: `${JSON.stringify(log)}`, error, isValid: false};
  }
  return result;
};


const hello: ValidatedEventAPIGatewayProxyEvent<typeof schema> = async (event) => {
  const body = JSON.parse(JSON.stringify(event)).body
  const idtoken = JSON.parse(JSON.stringify(event)).token
  const token : ClaimVerifyRequest = { token : idtoken}
  const auth = await handler(token);
  let validUser = true
  let returnMessage : string
  const testMessage = JSON.stringify(auth)
  if (auth.isValid) {
    validUser = true
    returnMessage = `${auth.userName}, says ${body.toUpperCase()}`
  }else{
    validUser = false
    returnMessage = `Could not authenticate the user`
  }

  return formatJSONResponse({
    message: returnMessage as string,
    validUser: validUser as boolean,
    event,
  });
}

export const main = middyfy(hello);
