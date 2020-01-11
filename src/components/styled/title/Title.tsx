import * as React from 'react';
import { StyledTitle, StyledH1, StyledH3 } from './TitleStyles';
import { BtnLink } from '../Buttons';

interface P {
  mainTitle: string;
  subTitle?: string;
  flex?: boolean;
  className?: string;
  light?: boolean;
  bgShadow?: boolean;
  cta?: boolean;
  ctaText?: string;
  linkTo?: string;
}

const Title: React.FC<P> = ({
  mainTitle,
  subTitle,
  flex,
  light,
  bgShadow,
  className,
  cta,
  ctaText,
  linkTo,
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
      {cta && (
        <BtnLink link to={linkTo?.toString()}>
          {ctaText}
        </BtnLink>
      )}
    </StyledTitle>
  );
};
export default Title;
