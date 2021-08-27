import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {UserModel} from '../models/user.model';
import {ToastrService} from 'ngx-toastr';

@Injectable({
  providedIn: 'root'
})
export class SettingsService {

  constructor(private http: HttpClient, private toastr: ToastrService) {}

  public loadMyProfile(): Observable<UserModel> {
    return this.http.get<UserModel>('http://localhost:8080/api/user/loadMyProfile');
  }

  public updateNotificationSettings(comment: boolean, message: boolean, post: boolean): void {
    this.http.post('http://localhost:8080/api/user/updateNotificationSettings', {
      receiveCommentNotifications: comment,
      receiveMessagesNotifications: message,
      receivePostNotifications: post
    }).subscribe(
        res => {
          this.toastr.success('Notification settings updated!');
        },
        (error => {
          console.log(error);
          this.toastr.error('Update failed');
        })
      );
  }

  updatePrivacySettings(privacy: boolean, messages: boolean, tag: boolean): void {
    this.http.post('http://localhost:8080/api/user/updatePrivacySettings', {
      isPrivate: privacy,
      receiveMessages: messages,
      taggable: tag
    }).subscribe(
      res => {
        this.toastr.success('Privacy settings updated!');
      },
      (error => {
        console.log(error);
        this.toastr.error('Update failed');
      })
    );
  }

  updateProfile(firstName: string, lastName: string, Email: string, PhoneNumber: string, web: string, Bio: string,
                Gender: string, Birth: Date): void {
    this.http.post('http://localhost:8080/api/user/updateProfileInfo', {
      email: Email,
      name: firstName,
      surname: lastName,
      phoneNumber: PhoneNumber,
      gender: Gender,
      birth: Birth,
      webSite: web,
      bio: Bio
    }).subscribe(
      res => {
        this.toastr.success('Profile updated!');
      },
      (error => {
        console.log(error);
        this.toastr.error('Update failed');
      })
    );
  }

  sendVerificationRequest(fd: FormData): void{
    this.http.post('http://localhost:8080/api/verification/user', fd)
      .subscribe(
        res => {
          this.toastr.success('Request sent!');
        },
        (error => {
          this.toastr.error('Request failed!');
        })
      );
  }
}

