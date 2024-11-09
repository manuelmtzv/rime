import type { User } from "./user.type";

export type LoginRequest = {
  identifier: string;
  password: string;
};

export type AuthResponse = {
  data: {
    user: User;
    token: string;
    refreshToken: string;
  };
};
