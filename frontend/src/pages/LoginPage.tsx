import { Wrapper } from "../components/Wrapper";
import { Title } from "../components/Title";

const LoginPage = () => {
  const loginGithub = () => {
    window.location.href = "/login?type=github";
  };

  const loginYandex = () => {
    window.location.href = "/login?type=yandex";
  };

  const loginByCode = () => {
    window.location.href = "/login?type=code";
  };

  return (
    <Wrapper>
      <Title>Log in</Title>

      <button onClick={loginGithub}>
        GitHub
      </button>

      <button onClick={loginYandex}>
        Yandex
      </button>

      <button onClick={loginByCode}>
        Code
      </button>
    </Wrapper>
  );
};

export default LoginPage;
