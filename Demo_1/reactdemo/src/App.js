import logo from './logo.svg';
import LoginButton from "./component/LoginButton"
import InputForm from "./component/Inputform"
import './App.css';

function App() {

  //First check that there is a token in the url
  const token = getToken()
  if (token == null){

    //Prompt user to log in to continue
    return (
      <div className="container">
        <header> Please log in to continue </header>
        <br/>
        <LoginButton/>
      </div>
    )
  }

  //The actual app starts here
  //Return only if the user already has a token
  return (
    <div className="container">
      <h1>Demo 1</h1>
      This demo takes some message as input and sends it to a lambda server to capitalize it.
      <br/>
      <br/>
      <InputForm token={token}/>

    </div>
  );
}


//Simply gets the token from the url, no authentication here
const getToken = () => {
  var url = document.location.hash;
  var access_token = new URLSearchParams(url).get(`access_token`);
 
  return access_token
}


export default App;
