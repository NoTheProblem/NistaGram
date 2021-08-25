export interface VerificationRequestModel {
  id: string;
  username: string;
  firstName: string;
  lastName: string;
  category: string;
  path: string;
  dateSubmitted: Date;
  dateAnswered: Date;
  answer: string;
  isAnswered: boolean;
  image: string;
}


