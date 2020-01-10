import * as React from 'react';
import { StyledTitle, StyledH1, StyledH3 } from './TitleStyles';

interface P {
  mainTitle: string;
  subTitle?: string;
  flex?: boolean;
  className?: string;
  light?: boolean;
  bgShadow?: boolean;
}

const Title: React.FC<P> = ({
  mainTitle,
  subTitle,
  flex,
  light,
  bgShadow,
  className,
}) => {
  return (
    <StyledTitle
      flex={flex}
      light={light}
      bgShadow={bgShadow}
      className={className}
    >
      <StyledH1> {mainTitle} </StyledH1>
      <StyledH3>{subTitle}</StyledH3>
    </StyledTitle>
  );
};
export default Title;
