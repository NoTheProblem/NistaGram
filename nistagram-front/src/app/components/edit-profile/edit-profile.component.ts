import {Component, Input, OnInit} from '@angular/core';
import {UserModel} from '../../models/user.model';
import {SettingsService} from '../../services/settings.service';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component..html',
  styleUrls: ['./edit-profile.component.css']
})
export class EditProfileComponent implements OnInit {
  @Input() userProfile: UserModel;
  firstName: string;
  lastName: string;
  email: string;
  phoneNumber: string;
  web: string;
  bio: string;
  gender: string;
  birth: Date;

  constructor(private settingsService: SettingsService) { }


  ngOnInit(): void {
    setTimeout(() => {
      this.firstName = this.userProfile.name;
      this.lastName = this.userProfile.surname;
      this.email = this.userProfile.email;
      this.phoneNumber = this.userProfile.phoneNumber;
      this.web = this.userProfile.webSite;
      this.bio = this.userProfile.bio;
      this.gender = this.userProfile.gender;
      this.birth = this.userProfile.birth;
      }, 1000);

  }

  UpdateProfile(): void {
    this.settingsService.updateProfile(this.firstName, this.lastName,this.email,  this.phoneNumber, this.web,
      this.bio, this.gender, this.birth);
  }
}
