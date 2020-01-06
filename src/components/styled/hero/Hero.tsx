import * as React from 'react';
import { StyledHero, Banner } from './HeroStyles';

interface Props {
  children?: React.ReactNode;
  title?: string;
  text?: string;
}

const Hero: React.FC<Props> = ({ children, title, text }) => {
  return (
    <>
      <StyledHero>
        <Banner>
          <h3>{title}</h3>
          <p>{text}</p>
        </Banner>
        {children}
      </StyledHero>
    </>
  );
};
export default Hero;
