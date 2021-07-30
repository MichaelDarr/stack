const getEnv = (varString: string): string | null => {
    const extractedVar = process.env[varString];
    if (!extractedVar) {
        return null;
    }
    return extractedVar;
};

const getNumericEnv = (varString: string): number | null => {
    const envVal = getEnv(varString);
    if (!envVal) {
        return null;
    }
    const numbericEnvVal = parseInt(envVal);
    if (isNaN(numbericEnvVal)) {
        throw new Error(
            `Value of numeric environment variable ${varString} ` +
            `(${envVal}) is not a number.`);
    }
    return numbericEnvVal;
};

const getRequiredEnv = (varString: string): string => {
    const envVal = getEnv(varString);
    if (!envVal) {
        throw new Error(
            `Required environment variable not found: ${varString}`);
    }
    return envVal;
};

const getRequiredNumericEnv = (varString: string): number => {
    const envVal = getNumericEnv(varString);
    if (!envVal) {
        throw new Error(
            `Required numeric environment variable not found: ${varString}`);
    }
    return envVal;
};

export const Env = {
    authHost: getRequiredEnv('AUTH_HOST'),
    authPort: getRequiredNumericEnv('AUTH_PORT'),
    authScheme: getRequiredEnv('AUTH_SCHEME'),
    backendHost: getRequiredEnv('BACKEND_HOST'),
    backendPort: getRequiredNumericEnv('BACKEND_PORT'),
    backendPath: getRequiredEnv('BACKEND_GQL_PATH'),
    backendScheme: getRequiredEnv('BACKEND_SCHEME'),
};

export const authURL = (
    `${Env.authScheme}://${Env.authHost}:${Env.authPort}`
)

export const backendURL = (
    `${Env.backendScheme}://${Env.backendHost}:${Env.backendPort}${Env.backendPath}`
);
