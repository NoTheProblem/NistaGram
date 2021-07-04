import { Injectable } from '@angular/core';
import jwt_decode from 'jwt-decode';

const TOKEN_KEY = 'auth-token';
const USER_KEY = 'auth-user';

@Injectable({
  providedIn: 'root'
})
export class TokenStorageService {
  private user: any;
  private decoded: any;

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

  public getUser(): any {
    return JSON.parse(sessionStorage.getItem(USER_KEY));
  }

  public getUsername(): string{
    this.user = this.getToken();
    let sub: any;
    ({sub} = jwt_decode(this.user));
    return sub;
  }

  public getUserType(): string{
    this.user = this.getToken();
    this.decoded = jwt_decode(this.user);
    return this.decoded.type;
  }

  public getUserID(): number{
    this.user = this.getToken();
    this.decoded = jwt_decode(this.user);
    return this.decoded.id;
  }

  public permissionForPage(role: string): boolean{
    if (!this.isLoggedIn()) {
      return false;
    }
    if(this.getToken() !== role){
      return false;
    }
    return true;
  }

  public isLoggedIn(): boolean{
    this.user = this.getToken();
    if (this.user === null){
      return false;
    }
    return  true;
  }
}
