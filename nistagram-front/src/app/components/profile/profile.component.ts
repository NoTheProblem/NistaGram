import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {PostModel} from '../../models/post.model';
import {UserModel} from '../../models/user.model';
import {UserService} from '../../services/user.service';
import {TokenStorageService} from '../../_services/token-storage.service';
import {FollowService} from '../../services/follow.service';
import {RELATION_TYPE, RelationType} from '../../models/relationshipType.model';

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
  isError = false;
  isLogged = false;
  showFollowButton = true;
  isBlocked = false;
  relationType: RELATION_TYPE;

  constructor(private route: ActivatedRoute, private userService: UserService, private followService: FollowService,
              private tokenStorageService: TokenStorageService) { }

  ngOnInit(): void {
    const routeParam = this.route.snapshot.paramMap;
    this.username = routeParam.get('username');
    this.followService.getRelationship(this.username).subscribe((relationType: RelationType) => {
        this.relationType = relationType.relation;
      },
      error => {
        this.relationType = RELATION_TYPE.NOT_FOLLOWING;
      }
    );

    this.userService.loadProfile(this.username)
      .subscribe((profile: UserModel) => {
        this.profile = profile;
      },
        (error => {
          this.isError = true;
          console.log(error.error);
          this.error = error.error.slice(0, -5);
          if (this.error === 'record not found'){
            this.showFollowButton = false;
            this.error = 'User with that username does not exist!';
          }
        })
      );

    this.userService.loadProfilePosts(this.username)
      .subscribe((postsList: Array<PostModel>) => {
          this.posts = postsList;
        },
        (error => {
          this.isError = true;
          console.log(error.error);
          this.error = this.error + error.error;
        })
      );

    this.isLogged = this.tokenStorageService.isLoggedIn();
    if (this.isLogged){
      if ( this.username === this.tokenStorageService.getUsername()){
        this.isMe = true;
      }
    }
  }

  followUser(): void {
    this.followService.followUser(this.username);
  }

  unFollowUser(): void {
    this.followService.unFollowUser(this.username);
  }

  blockUser(): void {
    this.followService.blockUser(this.username);
  }

  unBlockUser(): void {
    this.followService.unBlockUser(this.username);
  }



}

