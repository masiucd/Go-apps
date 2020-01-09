import * as React from 'react';
import styled from 'styled-components';
import BackgroundImage, {
  IFluidObject,
  IBackgroundImageProps,
} from 'gatsby-background-image';
import { graphql, useStaticQuery } from 'gatsby';

const HeroData = graphql`
  query {
    heroBG: file(relativePath: { eq: "hero1.jpg" }) {
      childImageSharp {
        fluid(quality: 90, maxWidth: 4160) {
          ...GatsbyImageSharpFluid_withWebp
        }
      }
    }
  }
`;
interface Props {
  children?: React.ReactNode;
  className?: string;
  img: { src: string };
  main?: boolean;
  position?: boolean;
}

interface HeroDataProps {
  heroBG: {
    childImageSharp: {
      fluid?: IFluidObject;
    };
  };
}

const Hero: React.FC<Props> = ({
  children,
  className,
  img,
  main,
  position,
}) => {
  const data: HeroDataProps = useStaticQuery(HeroData);
  const { fluid } = data.heroBG.childImageSharp;

  return (
    <>
      <BackgroundImage className={className} fluid={img || fluid} main={main}>
        {children}
      </BackgroundImage>
    </>
  );
};

export default styled(Hero)`
  min-height: ${props => (props.main ? `calc(100vh - 84px)` : `50vh`)};
  background: ${props =>
      props.main
        ? `linear-gradient(rgba(191, 44, 29, 0.4), rgba(9, 1, 00, 0.7))`
        : `linear-gradient(rgba(11, 4, 29, 0.6), rgba(9, 1, 20, 0.2))`}fixed
    no-repeat;

  background-position: ${props => (props.position ? props.position : 'center')};
  background-size: cover;
  opacity: 1 !important;
  display: flex;
  justify-content: center;
  align-items: center;
`;
// <BackgroundImage
// Tag="section"
// className={className}
// fluid={imageData}
// backgroundColor={`#040e18`}
// ></BackgroundImage>
