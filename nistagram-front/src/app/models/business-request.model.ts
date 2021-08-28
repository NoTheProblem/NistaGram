export interface BusinessRequestModel {
  username: string;
  web: string;
  email: string;
  status: STATUS_TYPE;
}

export enum STATUS_TYPE  {
  PENDING,
  DECLINE ,
  ACCEPT
}
