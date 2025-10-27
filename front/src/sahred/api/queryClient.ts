import { QueryClient } from "@tanstack/react-query";
import type { AxiosError } from "axios";

export const queryClient = new QueryClient();

declare module "@tanstack/react-query" {
  interface Register {
    defaultError: AxiosError;
  }
}
