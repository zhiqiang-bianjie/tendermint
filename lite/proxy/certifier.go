package proxy

import (
	"github.com/tendermint/tendermint/lite"
	certclient "github.com/tendermint/tendermint/lite/client"
	"github.com/tendermint/tendermint/lite/files"
	"errors"
	liteErr "github.com/tendermint/tendermint/lite/errors"
)

func GetCertifier(chainID, rootDir, nodeAddr string) (*lite.InquiringCertifier, error) {
	trust := lite.NewCacheProvider(
		lite.NewMemStoreProvider(),
		files.NewProvider(rootDir),
	)

	source := certclient.NewHTTPProvider(nodeAddr)

	// XXX: total insecure hack to avoid `init`
	fc, err := source.LatestCommit()
	/* XXX
	// this gets the most recent verified commit
	fc, err := trust.LatestCommit()
	if certerr.IsCommitNotFoundErr(err) {
		return nil, errors.New("Please run init first to establish a root of trust")
	}*/
	if err != nil {
		return nil, err
	}

	cert, err := lite.NewInquiringCertifier(chainID, fc, trust, source)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func GetCertifierWithHeight(chainID, rootDir, nodeAddr string, height int64) (*lite.InquiringCertifier, error) {
	trust := lite.NewCacheProvider(
		lite.NewMemStoreProvider(),
		files.NewProvider(rootDir),
	)

	source := certclient.NewHTTPProvider(nodeAddr)
	//Load trusted commit from trust store
	fc, err := trust.LatestCommit()
	if !liteErr.IsCommitNotFoundErr(err) {
		return nil, errors.New(err.Error())
	}
	//If trust store is empty, get the commit at the specific height as the initial trusted commit
	fc, err = source.GetByHeight(height)
	if err != nil {
		return nil, err
	}

	cert, err := lite.NewInquiringCertifier(chainID, fc, trust, source)
	if err != nil {
		return nil, err
	}

	return cert, nil
}
