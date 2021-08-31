export interface UserModel {
  id: string;
  email: string;
  username: string;
  role: USER_TYPE;
  name: string;
  surname: string;
  phoneNumber: string;
  gender: string;
  birth: Date;
  webSite: string;
  bio: string;
  isPrivate: boolean;
  receiveMessages: boolean;
  taggable: boolean;
  receivePostNotifications: boolean;
  receiveCommentNotifications: boolean;
  receiveMessagesNotifications: boolean;
  numberOfPosts: number;
  numberOfFollowers: number;
  numberOfFollowing: number;
  verified: boolean;
}

export enum USER_TYPE  {
  REGULAR,
  ADMINISTRATOR,
  AGENT
}
