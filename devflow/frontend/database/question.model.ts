import Answer from "./answer.model";
import Tag from "./tag.model";
import User from "./user.model";

type Question = {
  _id: string;
  title: string;
  description: string;
  userID: string;
  user: User;
  tags: string[];
  tagDetails: Tag[];
  views: number;
  upvotes: string[];
  downvotes: string[];
  answers: Answer[]; //
  createdAt: Date; //
};

export default Question;
