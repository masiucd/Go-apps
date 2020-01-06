import * as React from 'react';
import uuid from 'uuid/v4';
import { Instagram, Linkedin, Github } from 'styled-icons/feather';

type SocialIcons = {
  id: string;
  icon: React.ReactNode;
  url: string;
};

const socialIcons: SocialIcons[] = [
  {
    id: uuid(),
    icon: <Instagram size="35" />,
    url: '',
  },
  {
    id: uuid(),
    icon: <Linkedin size="35" />,
    url: '',
  },
  {
    id: uuid(),
    icon: <Github size="35" />,
    url: '',
  },
];

export default socialIcons;
