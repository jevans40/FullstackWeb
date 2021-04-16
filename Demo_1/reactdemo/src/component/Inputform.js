import { setTextRange } from "typescript"
import {useState} from 'react'
import LoginButton from './LoginButton'

const InputForm = ({token}) => {
    const [text,setText] = useState('')
    const accessToken = token
    


    const onSubmit = (e) => {
        let response = {}
        e.preventDefault()
        if(!text){
            alert('Please enter some text')
            return
        }
        console.log(`https://7fzhl7q379.execute-api.us-east-2.amazonaws.com/dev/hello/${text}/${accessToken}/`)
        fetch(`https://7fzhl7q379.execute-api.us-east-2.amazonaws.com/dev/hello/${text}/${accessToken}/`, {method: 'GET'})
        .then(resp => resp.json())
        .then(data => {response = JSON.parse(data.body)})
        .then(() =>{
            //Async data here
            if (!response.validUser){
                alert("You are not authenticated please re-log in");
            }else{
                alert(response.message);
            }
        })
    }

    return (
        <form className='add-form' onSubmit={onSubmit}>
            <div className='form-Control'>
                <label>Input text here</label>
                <br/>
                <input type='text' 
                  placeholder='Message' 
                  value={text} 
                  onChange={(e)=> setText(e.target.value)} />
            </div>
            <button type="submit" className="btn btn-primary">Submit</button>
            <LoginButton/>
        </form>
    )
}

export default InputForm