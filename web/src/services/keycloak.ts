import Keycloak from 'keycloak-js';

let keycloakInstance: Keycloak;

export const initKeycloak = async (onAuthenticatedCallback: () => void) => {
  try {
    // Fetch the keycloak.json file from the public folder
    const response = await fetch('/keycloak.json');
    const keycloakConfig = await response.json();

    // Ensure the required properties are present
    if (!keycloakConfig['auth-server-url']) {
      throw new Error("The configuration object is missing the required 'auth-server-url' property.");
    }

    // Initialize Keycloak with the configuration from keycloak.json
    keycloakInstance = new Keycloak({
      realm: keycloakConfig.realm,
      url: keycloakConfig['auth-server-url'], // Use 'auth-server-url' as 'url'
      clientId: keycloakConfig.resource,
      publicClient: keycloakConfig['public-client'],
      sslRequired: keycloakConfig['ssl-required'],
      confidentialPort: keycloakConfig['confidential-port'],
    });

    await keycloakInstance.init({
      onLoad: 'login-required',
      redirectUri: window.location.origin, // Use the current origin as the redirect URI
    });

    if (keycloakInstance.authenticated) {
      console.log('Authenticated');
      onAuthenticatedCallback();
    } else {
      console.warn('Not authenticated');
    }
  } catch (error) {
    console.error('Failed to initialize Keycloak', error);
  }
};

export const getKeycloakInstance = () => keycloakInstance;
