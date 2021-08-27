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
  isError = false;
  err = '';

  constructor(private tokenStorageService: TokenStorageService, private followService: FollowService) { }

  ngOnInit(): void {
    this.isLogged = this.tokenStorageService.isLoggedIn();
    if (this.isLogged){
      this.followService.getRecommendedUsers().subscribe((users: string[]) => {
        this.usernames = users;
        if (this.usernames.length === 0 ){
          this.err = 'No suggestions available';
          this.isError = true;
        }
        },
        (error => {
          this.err = 'No suggestions.\nThere is some issue please come later';
          this.isError = true;
          console.log(error);
        })
      );
    }
  }

}
