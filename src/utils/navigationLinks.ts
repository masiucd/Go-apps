import uuid from 'uuid/v4';

type NavigationLinksType = {
  id: string;
  text: string;
  url: string;
};

const navLinks: NavigationLinksType[] = [
  {
    id: uuid(),
    text: 'home',
    url: '/',
  },
  {
    id: uuid(),
    text: 'about',
    url: '/about',
  },
  {
    id: uuid(),
    text: 'menu',
    url: '/menu',
  },
  {
    id: uuid(),
    text: 'contact',
    url: '/contact',
  },
];

export default navLinks;
