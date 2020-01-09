import * as React from 'react';
import { StyledTitle, StyledH1, StyledH3 } from './TitleStyles';

interface P {
  mainTitle: string;
  subTitle?: string;
  flex?: boolean;
  light?: boolean;
  bgShadow?: boolean;
}

const Title: React.FC<P> = ({ mainTitle, subTitle, flex, light, bgShadow }) => {
  return (
    <StyledTitle flex={flex} light={light} bgShadow={bgShadow}>
      <StyledH1> {mainTitle} </StyledH1>
      <StyledH3>{subTitle}</StyledH3>
    </StyledTitle>
  );
};
export default Title;
