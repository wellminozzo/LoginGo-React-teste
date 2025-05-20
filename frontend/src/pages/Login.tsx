import { useEffect, useState } from "react";
import axios from "axios";

function LoginPage() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

     useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) {
            window.location.href = "/boasvindas"; // Redireciona se já estiver logado
        }
    }, []);


    const handleLogin = async () => {
        try {
            const response = await axios.post<{ token: string; message: string }>("http://localhost:5000/login", {
                email,
                password,
            });
            const token = response.data.token;
            localStorage.setItem("token", token);
            alert(response.data.message);
            window.location.href = "/boasvindas";
        } catch (error) {
            if ((error as any).response) {
                alert(((error as any).response.data as { error: string }).error);
            } else {
                alert("Erro ao conectar com o servidor.");
            }
        }
    };

    return (
        <div>
            <h2>Login</h2>
            <input type="text" placeholder="Email" onChange={(e) => setEmail(e.target.value)} />
            <input type="password" placeholder="Senha" onChange={(e) => setPassword(e.target.value)} />
            <button onClick={handleLogin}>Entrar</button>
        </div>
    );
}

export default LoginPage;