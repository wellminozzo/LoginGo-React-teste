import axios from "axios";

function BoasVindas() {
    const fetchProtectedData = async () => {
        const token = localStorage.getItem("token");

        try {
            const response = await axios.get("http://localhost:5000/protected", {
                headers: { Authorization: token },
            });

            alert((response.data as { message: string }).message);
        } catch (error) {
            alert("Acesso negado! Faça login novamente.");
        }
    };

    return (
        <div>
            <h2>Bem-vindo ao sistema!</h2>
            <button onClick={fetchProtectedData}>Acessar área protegida</button>
        </div>
    );
}


export default BoasVindas;