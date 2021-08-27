import {Component, Input, OnInit} from '@angular/core';
import {PostModel} from '../../models/post.model';
import {UserModel} from '../../models/user.model';
import {SettingsService} from '../../services/settings.service';

@Component({
  selector: 'app-privacy-settings',
  templateUrl: './privacy-settings.component.html',
  styleUrls: ['./privacy-settings.component.css']
})
export class PrivacySettingsComponent implements OnInit {
  @Input() userProfile: UserModel;
  isPrivate: boolean;
  messages: boolean;
  tag: boolean;

  constructor(private settingsService: SettingsService) { }

  ngOnInit(): void {
    console.log(this.userProfile.isPrivate);
    this.isPrivate = this.userProfile.isPrivate;
    this.messages = this.userProfile.receiveMessages;
    this.tag = this.userProfile.taggable;
  }

  updateProfile(): void {
    console.log(this.isPrivate);
    this.settingsService.updatePrivacySettings(this.isPrivate, this.messages, this.tag);

  }
}
