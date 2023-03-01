import {Component} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'signup-component',
  templateUrl: 'signup.component.html',
  styleUrls: ['signup.component.scss'],
})
export class SignupComponent {
  email!: string;
  password!: string;
  password1!: string;

  constructor(private http: HttpClient) {}

  createUser() {
    if ((this.password == this.password1) && (this.email != null || this.password != null)){
      const params = { Email: this.email, Password: this.password }

      const httpOptions = {
        headers: new HttpHeaders({
          'Content-Type': 'application/json'
        })
      };

      this.http.post("http://localhost:8181/createUser", params, httpOptions)
      .subscribe(response => {
        console.log(response);
      });
    } else if (this.email == null || this.password == null){
      console.log("Email or Password is empty!");
    } else {
      console.log("Passwords don't match!");
    }
    
  }
}
