import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  constructor() { }

  ngOnInit(): void {
  }

  login(form: any): void {
    const {email, password} = form.value;


    if (!form.valid) {
      return;
    }
    // this.authService.SignIn(email, password);
    form.resetForm();
  }


}
