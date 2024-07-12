package main

import (
	"Go_HSE_2024/2_and_3_HW_server/proto"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"strconv"
	"time"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Bank",
		Short: "Create account",
	}
	host string
	port int
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var createAccountCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		amount, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Printf("invalid amount: %s", err)
			return
		}

		conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		defer func() {
			_ = conn.Close()
		}()

		cli := proto.NewBankClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		rep, err := cli.CreateAccount(ctx, &proto.CreateAccountRequest{Name: name, Amount: amount})
		if err != nil {
			panic(err)
		}
		fmt.Println(rep.Message)
	},
}

var getAccountCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Show account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		defer func() {
			_ = conn.Close()
		}()

		cli := proto.NewBankClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		rep, err := cli.GetAccount(ctx, &proto.GetAccountRequest{Name: name})
		if err != nil {
			panic(err)
		}
		fmt.Println(rep.Message)
	},
}

var deleteAccountCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete an account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		defer func() {
			_ = conn.Close()
		}()

		cli := proto.NewBankClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		rep, err := cli.DeleteAccount(ctx, &proto.DeleteAccountRequest{Name: name})
		if err != nil {
			panic(err)
		}
		fmt.Println(rep.Message)
	},
}

var patchCmd = &cobra.Command{
	Use:   "patch [name] [amount]",
	Short: "New amount on the account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		amount, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println("Invalid amount")
			return
		}

		conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		defer func() {
			_ = conn.Close()
		}()

		cli := proto.NewBankClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		rep, err := cli.Patch(ctx, &proto.PatchRequest{Name: name, Amount: amount})
		if err != nil {
			panic(err)
		}
		fmt.Println(rep.Message)
	},
}

var updateNameCmd = &cobra.Command{
	Use:   "new_name [name] [new name]",
	Short: "New name of account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		newName := args[1]
		conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		defer func() {
			_ = conn.Close()
		}()

		cli := proto.NewBankClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		rep, err := cli.UpdateName(ctx, &proto.UpdateNameRequest{Name: name, NewName: newName})
		if err != nil {
			panic(err)
		}
		fmt.Println(rep.Message)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&host, "host", "0.0.0.0", "host to bind to")
	rootCmd.PersistentFlags().IntVar(&port, "port", 5678, "port to bind to")
	rootCmd.AddCommand(getAccountCmd)
	rootCmd.AddCommand(createAccountCmd)
	rootCmd.AddCommand(deleteAccountCmd)
	rootCmd.AddCommand(patchCmd)
	rootCmd.AddCommand(updateNameCmd)
}
