import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {PostModel} from '../../models/post.model';
import {UserModel} from '../../models/user.model';
import {UserService} from '../../services/user.service';
import {ToastrService} from 'ngx-toastr';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  public profile: UserModel;
  public posts: Array<PostModel> = new Array<PostModel>();
  isLocked: boolean;
  public username: string;
  public error: string;

  constructor(private route: ActivatedRoute, private userService: UserService, private toastr: ToastrService) { }

  ngOnInit(): void {
    this.error = ""
    this.isLocked = false;
    const routeParam = this.route.snapshot.paramMap;
    this.username = routeParam.get('username');
    this.userService.loadProfile(this.username)
      .subscribe((profile: UserModel) => {
        this.profile = profile;
        console.log(profile);
      },
        (error => {
          this.error = error.error;
          this.isLocked = true;
          console.log(error);
          this.toastr.error(error.error);
        })
      );
    this.userService.loadProfilePosts(this.username)
      .subscribe((postsList: Array<PostModel>) => {
        this.posts = postsList;
      },
        (error => {
          console.log(error);
          this.toastr.error('Posts failed');
        })
      );

  }
}

