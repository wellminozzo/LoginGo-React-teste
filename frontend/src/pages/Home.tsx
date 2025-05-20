import React from 'react';
import { Link } from 'react-router-dom';

const Home: React.FC = () => {
    return (
        <div>
            <h1>Welcome to the Home Page</h1>
            <p>This is the main landing page of the application.</p>
            <ul>
                <li>Feature 1</li>
                <li>Feature 2</li>
                <li>Feature 3</li>
            </ul>
            <div style={{ marginTop: 24 }}>
                <Link to="/login" style={{ marginRight: 16 }}>Login</Link>
                <Link to="/cadastro">Cadastro</Link>
            </div>
        </div>
    );
};

export default Home;