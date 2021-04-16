import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router'

@Component({
  selector: 'app-input-form',
  templateUrl: './input-form.component.html',
  styleUrls: ['./input-form.component.css']
})


//This class currently just handles sending information to lambda and authenticating the user.
//If the user has an invalid token they must re authenticate before continuing
export class InputFormComponent implements OnInit {
  message : 'test';
  constructor(private router: Router) { }

  ngOnInit(): void {
  }
  

  //When the submit button is clicked
  onSubmit() {
    var url = document.location.hash;
    var access_token = new URLSearchParams(url).get(`access_token`);
 
    //Make sure text isnt empty
    if(!this.message){
      alert('Please enter some Text')
      return
    }


    //Fetch lambda api
    fetch(`https://7fzhl7q379.execute-api.us-east-2.amazonaws.com/dev/hello/${this.message}/${access_token}/`, {mode: 'cors', method: 'GET'})
    .then(resp => resp.json())
    .then(data => {let response = JSON.parse(data.body)
        if (!response.validUser){
            alert("You are not authenticated please re-log");
            this.router.navigateByUrl('/login');
        }else{
            alert(response.message);
        }
    })
    
  }
}
