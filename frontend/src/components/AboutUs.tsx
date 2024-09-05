import React from "react";

const AboutUs: React.FC = () => (
  <div className="p-8 pt-20 rounded-lg max-w-3xl mx-auto">
    <h1 className="text-4xl font-bold text-center mb-6">About Us</h1>
    <p className="text-lg leading-relaxed mb-4 text-center">
      <span className="font-bold text-green-700">Nexos</span> is an Arch
      Linux-based operating system designed specifically for the Nexus Security
      Club. It aims to provide a user-friendly experience for individuals with
      any skill level by including all necessary tools and features out of the
      box.
    </p>
    <p className="text-lg leading-relaxed mb-4 text-center">
      <span className="font-bold text-green-700">Nexos</span> is designed to be
      user-friendly for individuals with varying levels of skill. Everything
      needed should be provided right from the start.
    </p>
  </div>
);

export default AboutUs;
