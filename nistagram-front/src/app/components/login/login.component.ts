import {Component, OnInit} from '@angular/core';
import {AuthService} from '../../_services/auth.service';
import {TokenStorageService} from '../../_services/token-storage.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  roles = '';
  username = '';
  isLoggedIn = false;
  isLoginFailed = false;
  errorMessage = '';
  private user: any;
  constructor(private authService: AuthService, private tokenStorage: TokenStorageService) { }

  ngOnInit(): void {
  }

  login(form: any): void {
    const {email, password} = form.value;
    console.log(email);
    console.log(password);
    this.authService.login(form.username, form.pass).subscribe(
      data => {
        this.tokenStorage.saveToken(data.accessToken);
        this.tokenStorage.saveUser(data);
        this.user = this.tokenStorage.getToken();
        this.isLoginFailed = false;
        this.isLoggedIn = true;
        this.username = this.tokenStorage.getUsername();
        this.reloadPage();
        this.roles = this.tokenStorage.getUserType();

      },
      err => {
        this.errorMessage = err.error.message;
        this.isLoginFailed = true;
      }
    );
  }

  reloadPage(): void {
    window.location.reload();
    this.roles = this.tokenStorage.getUserType();
  }


}

