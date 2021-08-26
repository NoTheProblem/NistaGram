import { Component, OnInit } from '@angular/core';
import {TokenStorageService} from '../../_services/token-storage.service';
import { Router } from '@angular/router';

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
  search: '';

  constructor(private tokenStorageService: TokenStorageService, private router: Router) {

  }

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

  searchQuery(): void {
    this.router.navigate(['search'], { state: { searchQuery: this.search }});
  }

}
