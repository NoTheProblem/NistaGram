import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {PostModel} from '../../models/post.model';
import {UserModel} from '../../models/user.model';
import {UserService} from '../../services/user.service';
import {TokenStorageService} from '../../_services/token-storage.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  public profile: UserModel;
  public posts: Array<PostModel> = new Array<PostModel>();
  isLocked: boolean;
  isLogged: boolean;
  isFollowed: boolean;
  public username: string;
  public error: string;
  public isMe: boolean;

  constructor(private route: ActivatedRoute, private userService: UserService, private tokenStorageService: TokenStorageService) { }

  ngOnInit(): void {
    this.isFollowed = true;
    this.error = '';
    this.isLogged = false;
    this.isMe = false;
    this.isLocked = false;
    const routeParam = this.route.snapshot.paramMap;
    this.username = routeParam.get('username');
    if (this.tokenStorageService.isLoggedIn()){
      this.isLogged = true;
      if (this.tokenStorageService.getUsername() === this.username){
        this.isMe = true;
        this.isFollowed = true;
      }
    }
    this.userService.loadProfile(this.username)
      .subscribe((profile: UserModel) => {
        this.profile = profile;
      },
        (error => {
          this.error = error.error.slice(0, -5);
          if (this.error === 'record not found'){
            this.error = 'User with that username does not exist!';
          }
          this.isLocked = true;
        })
      );

    // TODO call followers to see if he is followed if its locked

    this.userService.loadProfilePosts(this.username)
      .subscribe((postsList: Array<PostModel>) => {
          this.posts = postsList;
        },
        (error => {
          console.log(error);
        })
      );
  }
}

