import {useParams, useNavigate} from "react-router-dom";
import {API_URL} from './config';
import {useState, useEffect, useRef} from 'react';

import Hat from './Hat';
import './MainPage.css';
import './ProductPage.css';

const ProductCard = ({
                         initialProductID,
                         initialProductImageUrl,
                         productName,
                         productDescription,
                         productPrice,
                         productCount
                     }) => {
    const [productData, setProductData] = useState({
        productName: productName,
        productDescription: productDescription,
        productImageUrl: initialProductImageUrl,
        productPrice: productPrice,
        productCount: productCount,
    });

    const [productID, setProductID] = useState(initialProductID);
    const [productImageUrl, setProductImageUrl] = useState(initialProductImageUrl);

    const handleEdit = (field, value) => {
        console.log('[involodya]', field, value, productID)
        if (productData[field] === value) {
            return
        }

        const updatedProductData = {
            ...productData,
            [field]: value
        };

        setProductData(updatedProductData);

        let jsonData = JSON.stringify({
            ID: Number(productID),
            Count: Number(updatedProductData.productCount),
            Name: updatedProductData.productName,
            Description: updatedProductData.productDescription,
            ImageUrl: updatedProductData.productImageUrl,
            Cost: parseFloat(updatedProductData.productPrice)
        });

        (async () => {
            const token = window.localStorage.getItem("access");
            const response = await fetch(`${API_URL}/item/${productID}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
                body: jsonData,
            });
            const data = await response.json();
            setProductID(data.ID);
            console.log(data);
        })().catch((error) => {
            console.error('Error:', error);
        });
    };

    return (
        <div className="product">
            <div className="image-container">
                <img src={productImageUrl}
                     alt={productName} className="product-image"
                     onClick={() => {
                         const newUrl = prompt("Please enter the new image URL:");
                         if (newUrl) {
                             handleEdit('productImageUrl', newUrl);
                             setProductImageUrl(newUrl)
                         }
                     }}
                />
            </div>
            <div className="info-container">
                <div
                    className="product-name"
                    contentEditable
                    suppressContentEditableWarning
                    onBlur={(e) => handleEdit('productName', e.target.innerText)}
                >
                    {productName}
                </div>
                <div
                    className="product-description"
                    contentEditable
                    suppressContentEditableWarning
                    onBlur={(e) => handleEdit('productDescription', e.target.innerText)}
                >
                    {productDescription}
                </div>
                <div
                    className="product-price"
                    contentEditable
                    suppressContentEditableWarning
                    onBlur={(e) => handleEdit('productPrice', e.target.innerText.replace(/^\$/, ''))}
                >
                    ${productPrice}
                </div>
                <div className="product-count">
                    <button onClick={() => {
                        const newCount = Math.max(1, productData.productCount - 1);
                        handleEdit('productCount', newCount);
                    }}>-
                    </button>
                    <span>{productData.productCount}</span>
                    <button onClick={() => {
                        const newCount = productData.productCount + 1;
                        handleEdit('productCount', newCount);
                    }}>+
                    </button>
                </div>
            </div>
        </div>
    );
};

function ProductPage() {
    const navigate = useNavigate();

    const {id} = useParams();
    const [productID, setProductID] = useState(id);
    const [product, setProduct] = useState({
        ImageUrl: '',
        Name: 'Add name',
        Description: 'Add description',
        Cost: 0,
        Count: 1
    });
    const [isLoading, setIsLoading] = useState(true);
    const hasCreatedItem = useRef(false);

    useEffect(() => {
        if (productID) {
            (async () => {

                const token = window.localStorage.getItem("access");
                const response = await fetch(`${API_URL}/item/${productID}`, {
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                });
                if (response.status === 403) {
                    navigate('/product/new');
                    window.location.reload();
                } else {
                    const data = await response.json();
                    setProduct(data);
                    setProductID(data.ID);
                    setIsLoading(false);
                }
            })();
        } else if (!id && !hasCreatedItem.current) {
            hasCreatedItem.current = true;
            (async () => {
                const token = window.localStorage.getItem("access");
                const response = await fetch(`${API_URL}/item/create`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    },
                    body: JSON.stringify(product),
                });
                const data = await response.json();
                setProduct(data);
                setProductID(data.ID);
                setIsLoading(false);
            })();
        }

    }, [productID]);

    const renderProductCards = () => (
        <ProductCard
            initialProductID={productID}
            initialProductImageUrl={product.ImageUrl}
            productName={product.Name}
            productDescription={product.Description}
            productPrice={product.Cost}
            productCount={product.Count}
        />
    );

    const renderCard = () => (
        <section className="content">
            {isLoading ? <div>Loading...</div> : renderProductCards()}
        </section>
    );

    return (
        <main className="main-page">
            <Hat/>
            {renderCard()}
        </main>
    );
}

export default ProductPage;