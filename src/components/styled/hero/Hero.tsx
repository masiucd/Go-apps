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
}

interface HeroDataProps {
  heroBG: {
    childImageSharp: {
      fluid?: IFluidObject;
    };
  };
}

const Hero: React.FC<Props> = ({ children, className, img, main }) => {
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
  min-height: ${props => (props.main ? `calc(100vh - 62px)` : `50vh`)};
`;
// <BackgroundImage
// Tag="section"
// className={className}
// fluid={imageData}
// backgroundColor={`#040e18`}
// ></BackgroundImage>
