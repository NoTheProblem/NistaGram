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
  searchParam: string;

  constructor(private tokenStorageService: TokenStorageService, private router: Router) { }

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
    switch (this.search[0]){
      case '@':
        this.searchParam = 'SearchAddress';
        console.log('SearchAddress');
        break;
      case '#':
        this.searchParam = 'SearchTag';
        console.log('SearchTag');
        break;
      default:
        this.searchParam = 'SearchUser';
        console.log('SearchUser');
    }
    this.router.navigate(['search']);
    console.log(this.search);
  }

}
