import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  public profile: ProfileModel;
  public posts: Array<Post> = new Array<Post>();
  public post: Post;
  public link = 'https://images.unsplash.com/photo-1511765224389-37f0e77cf0eb?w=500&h=500&fit=crop';

  constructor() { }

  ngOnInit(): void {
    this.profile = new ProfileModel('slavBrain', 'bijografija', 'Branislav', 'Lazarevic', 120, 100,20)
    this.post = new Post('https://images.unsplash.com/photo-1511765224389-37f0e77cf0eb?w=500&h=500&fit=crop', 10, 1, false);
    this.posts.push(this.post);
    this.post = new Post('https://images.unsplash.com/photo-1497445462247-4330a224fdb1?w=500&h=500&fit=crop', 20, 2, true);
    this.posts.push(this.post);
    this.post = new Post('https://images.unsplash.com/photo-1502630859934-b3b41d18206c?w=500&h=500&fit=crop', 30, 3, false);
    this.posts.push(this.post);
  }

}

export class ProfileModel{
  constructor(
    public username: string,
    public bio: string,
    public name: string,
    public surname: string,
    public numberOfPosts: number,
    public numberOfFollowers: number,
    public numberOfFolloweing: number
  ) {
  }
}

export class Post{
  constructor(
    public url: string,
    public likes: number,
    public comments: number,
    public collection: boolean
  ) {
  }
}


