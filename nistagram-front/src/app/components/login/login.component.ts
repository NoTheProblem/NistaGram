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
  isLoggedIn: boolean;
  isLoginFailed = false;

  constructor(private authService: AuthService, private tokenStorage: TokenStorageService) { }

  ngOnInit(): void {
    if (this.tokenStorage.isLoggedIn()){
      this.isLoggedIn = true;
      this.username = this.tokenStorage.getUsername();
    }
    else{
      this.isLoggedIn = false;
    }
  }

  login(form: any): void {
    const {email, password} = form.value;
    this.authService.login(email, password).subscribe(
      data => {
        this.tokenStorage.saveToken(data.token);
        this.tokenStorage.saveUser(data);
        window.location.reload();
      },
      err => {
        this.isLoginFailed = true;
      }
    );
  }

}

