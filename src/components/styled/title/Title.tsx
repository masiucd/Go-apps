import * as React from 'react';
import { StyledTitle, StyledH1, StyledH3 } from './TitleStyles';

interface P {
  mainTitle: string;
  subTitle?: string;
  flex?: boolean;
  light?: boolean;
}

const Title: React.FC<P> = ({ mainTitle, subTitle, flex, light }) => {
  return (
    <StyledTitle flex={flex} light={light}>
      <StyledH1> {mainTitle} </StyledH1>
      <StyledH3>{subTitle}</StyledH3>
    </StyledTitle>
  );
};
export default Title;
