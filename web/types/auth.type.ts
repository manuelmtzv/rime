import type { User } from "./user.type";

export type LoginRequest = {
  username: string;
  password: string;
};

export type RegisterRequest = {
  name: string;
  lastname: string;
  username: string;
  email: string;
  password: string;
};

export type AuthResponse = {
  data: {
    user: User;
    token: string;
    refreshToken: string;
  };
};
