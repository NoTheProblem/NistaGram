import { Component, OnInit } from '@angular/core';
import {PostService} from '../../services/post.service';

@Component({
  selector: 'app-upload-post',
  templateUrl: './upload-post.component.html',
  styleUrls: ['./upload-post.component.css']
})
export class UploadPostComponent implements OnInit {

  selectedFile: File = null;
  isUploaded = false;
  fileName = '';
  description: string;
  location: string;
  tag: string;
  addedTags: string[];
  addedTagsShow: string[];

  constructor(private postService: PostService) { }

  ngOnInit(): void {
    this.addedTagsShow = [];
    this.addedTags  = [];
  }

  onFileSelected(event): void{
    // tslint:disable-next-line:prefer-const
    // let reader = new FileReader();
    this.selectedFile = event.target.files[0];
    this.isUploaded = true;
    this.fileName = this.selectedFile.name;
    console.log(event.target.files);
  }

  addTag(newTag: string): void{
    this.addedTags.push(newTag);
    newTag = '#' + newTag;
    this.addedTagsShow.push(newTag);
    this.tag = null;
  }

  uploadPost(): void {
    const fd = new FormData();
    fd.append('myFile', this.selectedFile, this.selectedFile.name);
    fd.append('description', this.description);
    fd.append('location', this.location);
    fd.append('tags',  JSON.stringify(this.addedTags));
    this.postService.uploadPost(fd);
    this.fileName = null;
    this.isUploaded = false;
    this.description = null;
    this.addedTags = null;
    this.addedTagsShow = null;
    this.location = null;
  }



}
