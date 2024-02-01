import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './styles/ProfilePage.css'; // Make sure to create this CSS file for styling

const ProfilePage = () => {
  const navigate = useNavigate();
  const [userInfo, setUserInfo] = useState({
    firstname: '',
    lastname: '',
    phone: '',
    username: '',
    password: '',
    image: null,
    bio: '',
  });

  useEffect(() => {
    // Placeholder for fetching user info from backend
    // Set it in state, assuming you get this data from your backend
    setUserInfo({
      ...userInfo,
      firstname: 'John',
      lastname: 'Doe',
      phone: '123-456-7890',
      username: 'johndoe',
      bio: 'Lorem ipsum dolor sit amet...',
    });
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setUserInfo({ ...userInfo, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // API call to save changes
    console.log(userInfo);
  };

  return (
    <div className="profile-page">
      <h2>Edit Profile</h2>
      <form onSubmit={handleSubmit} className="profile-form">
        <div className="form-group">
          <label htmlFor="firstname">First Name</label>
          <input type="text" id="firstname" name="firstname" value={userInfo.firstname} onChange={handleChange} />
        </div>
        <div className="form-group">
          <label htmlFor="lastname">Last Name</label>
          <input type="text" id="lastname" name="lastname" value={userInfo.lastname} onChange={handleChange} />
        </div>
        {/* Repeat for other fields */}
        <div className="form-group">
          <label htmlFor="bio">Bio</label>
          <textarea id="bio" name="bio" value={userInfo.bio} onChange={handleChange} />
        </div>
        <div className="form-actions">
          <button type="submit">Save Changes</button>
          <button type="button" onClick={() => navigate('/main')}>Cancel</button>
        </div>
      </form>
    </div>
  );
};

export default ProfilePage;
