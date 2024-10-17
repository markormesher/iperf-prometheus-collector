type IperfTcpOutput = {
  error?: string;
  end: {
    sum_sent: {
      bytes: string;
      seconds: string;
    };
    sum_received: {
      bytes: string;
      seconds: string;
    };
  };
};

export { IperfTcpOutput };
