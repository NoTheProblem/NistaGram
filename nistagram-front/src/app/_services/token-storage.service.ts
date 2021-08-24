import { Injectable } from '@angular/core';
import jwt_decode from 'jwt-decode';

const TOKEN_KEY = 'auth-token';
const USER_KEY = 'auth-user';

@Injectable({
  providedIn: 'root'
})
export class TokenStorageService {
  private user: any;

  constructor() { }

  signOut(): void {
    window.sessionStorage.clear();
  }

  public saveToken(token: string): void {
    window.sessionStorage.removeItem(TOKEN_KEY);
    window.sessionStorage.setItem(TOKEN_KEY, token);
  }

  public getToken(): string {
    return sessionStorage.getItem(TOKEN_KEY);
  }

  public saveUser(user): void {
    window.sessionStorage.removeItem(USER_KEY);
    window.sessionStorage.setItem(USER_KEY, JSON.stringify(user));
  }

  public isLoggedIn(): boolean{
    this.user = this.getToken();
    if (this.user === null){
      return false;
    }
    return  true;
  }

  public getUsername(): string{
    const decoded: {role: number, userId: string} = jwt_decode(this.getToken());
    return decoded.userId;
  }

  public getRole(): number{
    const decoded: {role: number, userId: string} = jwt_decode(this.getToken());
    return decoded.role;
  }

}
