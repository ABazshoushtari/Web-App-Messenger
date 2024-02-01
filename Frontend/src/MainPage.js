import React from 'react';
import { useNavigate } from 'react-router-dom';
import './styles/MainPage.css'; // Ensure you have the corresponding CSS file

const MainPage = () => {
  const navigate = useNavigate();

  return (
    <div className="main-page">
      <header className="main-header">
        <div className="profile-icon" onClick={() => navigate('/profile')}>
          <span className="material-icons">account_circle</span>
        </div>
      </header>
      {/* Chat list and other main page content go here */}
    </div>
  );
};

export default MainPage;
