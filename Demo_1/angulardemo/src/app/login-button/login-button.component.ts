import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login-button',
  templateUrl: './login-button.component.html',
  styleUrls: ['./login-button.component.css']
})


//This class simply handles the login page, there is just a button 
//that sends the user to cognito for authentication
export class LoginButtonComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
