import { Component, OnInit } from '@angular/core';
import {AuthService} from '../../_services/auth.service';
import {UserService} from '../../services/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  isSuccessful = false;
  isSignUpFailed = false;
  errorMessage = '';
  registerBusiness = false;
  constructor(private authService: AuthService) { }

  ngOnInit(): void {
  }

  register(form: any): void {
    console.log(form.form.value)
    if (form.form.value.password !== form.form.value.passwordConf){
      this.isSignUpFailed = true;
      this.errorMessage = 'Passwords dont match!';
      return;
    }
    if (!this.registerBusiness){
      this.authService.register(form.form.value).subscribe(
        data => {
          this.isSuccessful = true;
          this.isSignUpFailed = false;
        },
        err => {
          this.isSignUpFailed = true;
        }
      );
    }
    else {
      this.authService.registerBusiness(form.form.value).subscribe(
        data => {
          this.isSuccessful = true;
          this.isSignUpFailed = false;
        },
        err => {
          this.isSignUpFailed = true;
        }
      );
    }

  }
}
