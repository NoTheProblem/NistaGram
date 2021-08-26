import { Component, OnInit } from '@angular/core';
import {TokenStorageService} from '../../_services/token-storage.service';
import {FollowService} from '../../services/follow.service';

@Component({
  selector: 'app-suggested',
  templateUrl: './suggested.component.html',
  styleUrls: ['./suggested.component.css']
})
export class SuggestedComponent implements OnInit {
  isLogged: boolean;
  usernames: string[];

  constructor(private tokenStorageService: TokenStorageService, private followService: FollowService) { }

  ngOnInit(): void {
    this.isLogged = this.tokenStorageService.isLoggedIn();
    if (this.isLogged){
      this.followService.getRecommendedUsers().subscribe((users: string[]) => {
        this.usernames = users;
        },
        (error => {
          this.usernames = [];
          this.usernames.push('slav');
          this.usernames.push('slav');
          this.usernames.push('slav');
          this.usernames.push('slav');
          this.usernames.push('slav');
          console.log(error);
        })
      );
    }
  }

}
