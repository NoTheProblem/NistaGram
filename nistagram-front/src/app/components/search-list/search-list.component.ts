import { Component, OnInit } from '@angular/core';
import {UserModel} from '../../models/user.model';
import {PostModel} from '../../models/post.model';
import {SearchService} from '../../services/search.service';

@Component({
  selector: 'app-search-list',
  templateUrl: './search-list.component.html',
  styleUrls: ['./search-list.component.css']
})
export class SearchListComponent implements OnInit {
  searchInput: string;
  isUsers = false;
  isPosts = false;
  users: Array<UserModel> = new Array<UserModel>();
  posts: Array<PostModel> = new Array<PostModel>();

  constructor(private searchService: SearchService) {

  }

  ngOnInit(): void {
    this.searchInput = '';
    if (history.state.searchQuery !== undefined){
      this.searchInput = history.state.searchQuery;
    }
    this.executeQuery();
  }

  executeQuery(): void{
    this.isUsers = false;
    this.isPosts = false;
    switch (this.searchInput[0]){
      case '#':
        this.isPosts = true;
        this.searchService.searchTag(this.searchInput.slice(1))
          .subscribe((postsList: Array<PostModel>) => {
            this.posts = postsList;
          });
        break;
      case '@':
        this.isPosts = true;
        this.searchService.searchLocation(this.searchInput.slice(1))
          .subscribe((postsList: Array<PostModel>) => {
            this.posts = postsList;
          });
        break;
      default:
        this.isUsers = true;
        this.searchService.searchUser(this.searchInput)
          .subscribe((usersList: Array<UserModel>) => {
            this.users = usersList;
          });
        break;
    }
  }

}
