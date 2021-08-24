import { Component, OnInit } from '@angular/core';
import {TokenStorageService} from '../../_services/token-storage.service';

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.css']
})
export class NavigationComponent implements OnInit {
  user: string;
  isLoggedIn = false;
  username: string;
  role: number;

  constructor(private tokenStorageService: TokenStorageService) { }

  ngOnInit(): void {
    this.isLoggedIn = this.tokenStorageService.isLoggedIn();
    if (this.isLoggedIn){
      this.username = this.tokenStorageService.getUsername();
      this.role = this.tokenStorageService.getRole();
    }
  }

  logout(): void {
    this.tokenStorageService.signOut();
    window.location.reload();
  }
}
