import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router'

@Component({
  selector: 'app-input-form',
  templateUrl: './input-form.component.html',
  styleUrls: ['./input-form.component.css']
})
export class InputFormComponent implements OnInit {
  message : 'test';
  constructor(private router: Router) { }

  ngOnInit(): void {
  }
  
  onSubmit() {
    var url = document.location.hash;
    var access_token = new URLSearchParams(url).get(`access_token`);
 
    if(!this.message){
      alert('Please enter some Text')
      return
    }
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
