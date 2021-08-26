import { Component, OnInit } from '@angular/core';
import {PostService} from '../../services/post.service';

@Component({
  selector: 'app-upload-post',
  templateUrl: './upload-post.component.html',
  styleUrls: ['./upload-post.component.css']
})
export class UploadPostComponent implements OnInit {

  selectedFiles: Array<File> = new Array<File>();
  description: string;
  location: string;
  tag: string;
  addedTags: string[];
  addedTagsShow: string[];
  images: Array<string> = new Array<string>();
  numberOfImages: number;

  constructor(private postService: PostService) { }

  ngOnInit(): void {
    this.addedTagsShow = [];
    this.addedTags  = [];
    this.numberOfImages = 0;
  }

  onFileSelected(event): void{
    const reader = new FileReader();
    reader.readAsDataURL(event.target.files[0]);
    this.numberOfImages = this.numberOfImages + 1;
    reader.onload = (e: any) => {
      this.images.push(e.target.result);
    };
    this.selectedFiles.push(event.target.files[0]);
  }

  addTag(newTag: string): void{
    this.addedTags.push(newTag);
    newTag = '#' + newTag;
    this.addedTagsShow.push(newTag);
    this.tag = null;
  }

  uploadPost(): void {
    const fd = new FormData();
    fd.append('numberOfImages', JSON.stringify(this.numberOfImages));
    for (let i = 0; i < this.numberOfImages ; i++){
       fd.append('myFile' + String(i), this.selectedFiles[i], this.selectedFiles[i].name);
    }
    fd.append('description', this.description);
    fd.append('location', this.location);
    fd.append('tags',  JSON.stringify(this.addedTags));
    // TODO hardkodirano true za post MOVE to followers service?
    fd.append('isPublic', JSON.stringify(true));
    this.postService.uploadPost(fd);
  }


  removeImage(image: string): void {
      this.numberOfImages = this.numberOfImages - 1;
      this.images = this.images.filter(img => {
        return (
          img !== image
        );
      });
  }
}





