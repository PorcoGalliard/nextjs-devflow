type User = {
  _id: string;
  clerkID: string;
  firstName: string;
  lastName: string;
  bio: string | null;
  picture: string | null;
  email: string;
  location: string | null;
  portfolioWebsite: string | null;
  isAdmin: boolean;
  reputation: number | null;
  saved: string[] | null;
  joinedAt: string;
};

export default User;