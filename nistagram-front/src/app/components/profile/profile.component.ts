import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {PostModel} from '../../models/post.model';
import {UserModel} from '../../models/user.model';
import {UserService} from '../../services/user.service';
import {TokenStorageService} from '../../_services/token-storage.service';
import {FollowService} from '../../services/follow.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  public profile: UserModel;
  public posts: Array<PostModel> = new Array<PostModel>();
  public username: string;
  public error = '';
  public isMe = false;
  isLocked = false;
  isLogged = false;
  isFollowing = false;

  constructor(private route: ActivatedRoute, private userService: UserService, private followService: FollowService,
              private tokenStorageService: TokenStorageService) { }

  ngOnInit(): void {
    const routeParam = this.route.snapshot.paramMap;
    this.username = routeParam.get('username');
    this.userService.loadProfile(this.username)
      .subscribe((profile: UserModel) => {
        this.profile = profile;
      },
        (error => {
          this.error = error.error.slice(0, -5);
          if (this.error === 'record not found'){
            this.error = 'User with that username does not exist!';
          }
        })
      );
    this.userService.loadProfilePosts(this.username)
      .subscribe((postsList: Array<PostModel>) => {
          this.posts = postsList;
        },
        (error => {
          this.error = this.error + error.error;
        })
      );
    this.isLogged = this.tokenStorageService.isLoggedIn();
  }

  followUser(): void {
    this.followService.followUser(this.profile.username);
  }

  unFollowUser(): void {
    this.followService.unFollowUser(this.profile.username);
  }

  blockUser(): void {
    this.followService.blockUser(this.profile.username);
  }

  unBlockUser(): void {
    this.followService.unBlockUser(this.profile.username);
  }



}

