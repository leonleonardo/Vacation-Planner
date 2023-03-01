import {Component} from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';

@Component({
  selector: 'login-component',
  templateUrl: 'login.component.html',
  styleUrls: ['login.component.scss'],
})
export class LoginComponent {
  // Initializing variables to save input given from user within login component
  email!: string;
  password!: string;

  // Injecting HttpClient for http requests
  constructor(private http: HttpClient) {}

  // Initializing function to be called on "sign in" button click
  loginUser() {
    // Initializing struct of parameters for request body
    const params = { Email: this.email, Password: this.password }

    // Initializing header for JSON file request body abilities
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    };

    // Make post request with appropriate URL, the user/pass given by the user and JSON capabilities
    this.http.post("http://localhost:8181/loginUser", params, httpOptions)
    .subscribe(response => {
        // Print request response to JS console
        console.log(response);
  });
}
}