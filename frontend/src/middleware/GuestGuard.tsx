import { Navigate } from "react-router-dom";

interface GuestGuardProps {
    children: JSX.Element;
}

export function GuestGuard({ children }: GuestGuardProps) {
    const token = localStorage.getItem("token");

    return token ? <Navigate to="/boasvindas" /> : children;
}