import {Component, OnInit} from '@angular/core';
import {SettingsService} from '../../services/settings.service';
@Component({
  selector: 'app-verification',
  templateUrl: './verification.component.html',
  styleUrls: ['./verification.component.css']
})
export class VerificationComponent implements OnInit {

  selectedFile: File = null;
  isUploaded = false;
  fileName = '';
  firstName: string;
  lastName: string;
  category: string;

  constructor(private settingsService: SettingsService) { }

  ngOnInit(): void {
  }

  onFileSelected(event): void{
    this.selectedFile = event.target.files[0];
    this.isUploaded = true;
    this.fileName = this.selectedFile.name;
    console.log(event.target.files);
  }

  SendVerificationRequest(): void {
    const fd = new FormData();
    fd.append('myFile', this.selectedFile, this.selectedFile.name);
    fd.append('firstName', this.firstName);
    fd.append('lastName', this.lastName);
    fd.append('category', this.category);
    this.settingsService.sendVerificationRequest(fd);
  }
}
